package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type targetRow struct {
	Repo       string
	SourceFile string
	Line       int
	Receiver   string
	FuncName   string
}

type datasetItem struct {
	RelPath   string `json:"relpath"`
	LineNo    int    `json:"lineno"`
	FocalFunc string `json:"focalfunc"`
	FuncBody  string `json:"funcbody"`
	FuncName  string `json:"funcname"`
}

var safeFileName = regexp.MustCompile(`[^A-Za-z0-9_]+`)

func main() {
	targetCSV := flag.String("target-csv", "", "path to rq1 target manifest csv")
	datasetsRoot := flag.String("datasets-root", "./datasets_aligned", "output datasets root")
	reposRoot := flag.String("repos-root", "", "path to ~/code/ase-dataset")
	projectsRoot := flag.String("projects-root", "./projects", "RATester projects workspace root")
	onlyRepo := flag.String("only-repo", "", "optional single repo filter")
	flag.Parse()

	if *targetCSV == "" || *reposRoot == "" {
		fmt.Fprintln(os.Stderr, "--target-csv and --repos-root are required")
		os.Exit(2)
	}

	rows, err := loadTargetRows(*targetCSV, *onlyRepo)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := os.MkdirAll(*datasetsRoot, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.MkdirAll(*projectsRoot, 0o755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	moduleCache := map[string]string{}
	repoCounts := map[string]int{}
	for _, row := range rows {
		if err := ensureProjectLink(*projectsRoot, *reposRoot, row.Repo); err != nil {
			fmt.Fprintf(os.Stderr, "ensure project link failed for %s: %v\n", row.Repo, err)
			os.Exit(1)
		}
		modulePath, ok := moduleCache[row.Repo]
		if !ok {
			modulePath, err = readModulePath(filepath.Join(*reposRoot, row.Repo, "go.mod"))
			if err != nil {
				fmt.Fprintf(os.Stderr, "read go.mod failed for %s: %v\n", row.Repo, err)
				os.Exit(1)
			}
			moduleCache[row.Repo] = modulePath
		}
		item, datasetDir, outName, err := buildDatasetItem(*reposRoot, row, modulePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "build dataset item failed for %s %s:%d %s: %v\n", row.Repo, row.SourceFile, row.Line, row.FuncName, err)
			os.Exit(1)
		}
		outDir := filepath.Join(*datasetsRoot, row.Repo, datasetDir)
		if err := os.MkdirAll(outDir, 0o755); err != nil {
			fmt.Fprintf(os.Stderr, "mkdir failed for %s: %v\n", outDir, err)
			os.Exit(1)
		}
		payload, err := json.MarshalIndent(item, "", "    ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "marshal failed for %s: %v\n", outName, err)
			os.Exit(1)
		}
		if err := os.WriteFile(filepath.Join(outDir, outName), append(payload, '\n'), 0o644); err != nil {
			fmt.Fprintf(os.Stderr, "write failed for %s/%s: %v\n", outDir, outName, err)
			os.Exit(1)
		}
		repoCounts[row.Repo]++
	}

	repos := make([]string, 0, len(repoCounts))
	for repo := range repoCounts {
		repos = append(repos, repo)
	}
	sort.Strings(repos)
	for _, repo := range repos {
		fmt.Printf("%s\t%d\n", repo, repoCounts[repo])
	}
}

func loadTargetRows(csvPath, onlyRepo string) ([]targetRow, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	header, err := reader.Read()
	if err != nil {
		return nil, err
	}
	index := map[string]int{}
	for i, name := range header {
		index[strings.TrimPrefix(name, "\ufeff")] = i
	}
	required := []string{"repo", "source_file", "line", "receiver", "func_name"}
	for _, key := range required {
		if _, ok := index[key]; !ok {
			return nil, fmt.Errorf("missing csv header %q", key)
		}
	}

	rows := make([]targetRow, 0, 1024)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		repo := record[index["repo"]]
		if onlyRepo != "" && repo != onlyRepo {
			continue
		}
		line, err := strconv.Atoi(record[index["line"]])
		if err != nil {
			return nil, fmt.Errorf("invalid line for %s: %w", repo, err)
		}
		rows = append(rows, targetRow{
			Repo:       repo,
			SourceFile: record[index["source_file"]],
			Line:       line,
			Receiver:   record[index["receiver"]],
			FuncName:   record[index["func_name"]],
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		a, b := rows[i], rows[j]
		if a.Repo != b.Repo {
			return a.Repo < b.Repo
		}
		if a.SourceFile != b.SourceFile {
			return a.SourceFile < b.SourceFile
		}
		if a.Line != b.Line {
			return a.Line < b.Line
		}
		if a.FuncName != b.FuncName {
			return a.FuncName < b.FuncName
		}
		return a.Receiver < b.Receiver
	})
	return rows, nil
}

func ensureProjectLink(projectsRoot, reposRoot, repo string) error {
	target := filepath.Join(reposRoot, repo)
	linkPath := filepath.Join(projectsRoot, repo)
	if _, err := os.Stat(target); err != nil {
		return fmt.Errorf("repo root missing: %w", err)
	}
	info, err := os.Lstat(linkPath)
	if err == nil {
		if info.Mode()&os.ModeSymlink == 0 {
			return fmt.Errorf("existing path is not symlink: %s", linkPath)
		}
		resolved, err := filepath.EvalSymlinks(linkPath)
		if err != nil {
			return err
		}
		if resolved == target {
			return nil
		}
		return fmt.Errorf("symlink points to %s, want %s", resolved, target)
	}
	if !os.IsNotExist(err) {
		return err
	}
	return os.Symlink(target, linkPath)
}

func readModulePath(goModPath string) (string, error) {
	content, err := os.ReadFile(goModPath)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(content), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(trimmed, "module ")), nil
		}
	}
	return "", fmt.Errorf("module path not found in %s", goModPath)
}

func buildDatasetItem(reposRoot string, row targetRow, modulePath string) (datasetItem, string, string, error) {
	repoRelativeFile, absPath, err := resolveSourcePath(reposRoot, row)
	if err != nil {
		return datasetItem{}, "", "", err
	}
	src, err := os.ReadFile(absPath)
	if err != nil {
		return datasetItem{}, "", "", err
	}
	fset := token.NewFileSet()
	fileNode, err := parser.ParseFile(fset, absPath, src, parser.ParseComments)
	if err != nil {
		return datasetItem{}, "", "", err
	}
	funcDecl, err := findFuncDecl(fset, fileNode, row)
	if err != nil {
		return datasetItem{}, "", "", err
	}
	start := fset.PositionFor(funcDecl.Pos(), false).Offset
	end := fset.PositionFor(funcDecl.End(), false).Offset
	lbrace := fset.PositionFor(funcDecl.Body.Lbrace, false).Offset
	if start < 0 || end > len(src) || lbrace > len(src) || lbrace < start {
		return datasetItem{}, "", "", fmt.Errorf("invalid offsets for %s", row.SourceFile)
	}
	funcBody := strings.TrimSpace(string(src[start:end]))
	focalFunc := strings.TrimSpace(string(src[start:lbrace]))
	packageRelativeFile := strings.TrimPrefix(repoRelativeFile, row.Repo+"/")
	repoRelativeDir := filepath.Dir(packageRelativeFile)
	if repoRelativeDir == "." {
		repoRelativeDir = ""
	}
	importPath := modulePath
	if repoRelativeDir != "" {
		importPath += "/" + filepath.ToSlash(repoRelativeDir)
	}
	datasetDir := sanitizeImportPath(importPath)
	fileName := buildDatasetFileName(row)
	return datasetItem{
		RelPath:   "./projects/" + repoRelativeFile,
		LineNo:    row.Line,
		FocalFunc: focalFunc,
		FuncBody:  funcBody,
		FuncName:  row.FuncName,
	}, datasetDir, fileName, nil
}

func findFuncDecl(fset *token.FileSet, fileNode *ast.File, row targetRow) (*ast.FuncDecl, error) {
	var fallback *ast.FuncDecl
	for _, decl := range fileNode.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		if fd.Name == nil || fd.Name.Name != row.FuncName {
			continue
		}
		if receiverString(fd) != row.Receiver {
			continue
		}
		if fset.PositionFor(fd.Pos(), false).Line == row.Line {
			return fd, nil
		}
		if fallback == nil {
			fallback = fd
		}
	}
	if fallback != nil {
		return fallback, nil
	}
	return nil, fmt.Errorf("function not found")
}

func receiverString(fd *ast.FuncDecl) string {
	if fd.Recv == nil || len(fd.Recv.List) == 0 {
		return ""
	}
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, token.NewFileSet(), fd.Recv.List[0].Type); err != nil {
		return ""
	}
	return buf.String()
}

func sanitizeImportPath(importPath string) string {
	replacer := strings.NewReplacer("/", "_", ".", "_")
	return replacer.Replace(importPath)
}

func sanitizeFileStem(name string) string {
	cleaned := safeFileName.ReplaceAllString(name, "_")
	cleaned = strings.Trim(cleaned, "_")
	if cleaned == "" {
		return "target"
	}
	return cleaned
}

func buildDatasetFileName(row targetRow) string {
	base := strings.TrimSuffix(filepath.Base(row.SourceFile), filepath.Ext(row.SourceFile))
	parts := []string{sanitizeFileStem(base)}
	if row.Receiver != "" {
		parts = append(parts, sanitizeFileStem(row.Receiver))
	}
	parts = append(parts, sanitizeFileStem(row.FuncName), fmt.Sprintf("L%d", row.Line))
	return strings.Join(parts, "__") + ".json"
}

func pathJoinSlash(parts ...string) string {
	trimmed := make([]string, 0, len(parts))
	for _, part := range parts {
		if part == "" {
			continue
		}
		trimmed = append(trimmed, strings.Trim(part, "/"))
	}
	return strings.Join(trimmed, "/")
}

func resolveSourcePath(reposRoot string, row targetRow) (string, string, error) {
	trimmed := strings.Trim(filepath.ToSlash(row.SourceFile), "/")
	candidates := []string{
		trimmed,
		pathJoinSlash(row.Repo, trimmed),
		pathJoinSlash(row.Repo, strings.TrimPrefix(trimmed, row.Repo+"/")),
	}
	seen := map[string]struct{}{}
	for _, rel := range candidates {
		if rel == "" {
			continue
		}
		if _, ok := seen[rel]; ok {
			continue
		}
		seen[rel] = struct{}{}
		abs := filepath.Join(reposRoot, filepath.FromSlash(rel))
		if _, err := os.Stat(abs); err == nil {
			return rel, abs, nil
		}
	}
	return "", "", fmt.Errorf("no existing source path matched %q", row.SourceFile)
}
