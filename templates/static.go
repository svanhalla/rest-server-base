// Package static makes it possible to embed files when distributing the application
package static

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Assets contains embedded files.
//
//go:embed *
var Assets embed.FS

// GetFile hämtar innehållet i en fil som en byte-slice, antingen från det inbäddade filsystemet eller från disk.
func GetFile(name string) ([]byte, error) {
	// Kontrollera om katalogen "static/" finns
	if directoryExists("static") {
		// Läs filen från disken
		return readFileFromDisk("static", name)
	}

	// Läs filen från inbäddade Assets
	return readFileFromAssets(name)
}

// directoryExists kontrollerar om en given katalog finns.
func directoryExists(path string) bool {
	info, err := os.Stat(path)

	return err == nil && info.IsDir()
}

func readFileFromDisk(basePath, filePath string) ([]byte, error) {
	// Sanera och validera filvägen
	cleanPath := filepath.Clean(filePath)
	if strings.HasPrefix(cleanPath, "..") || strings.HasPrefix(cleanPath, "/") {
		return nil, fmt.Errorf("invalid file path")
	}

	// Bygg den fullständiga sökvägen
	fullPath := filepath.Join(basePath, cleanPath)

	// Läs filen från den säkra sökvägen
	data, err := os.ReadFile(fullPath) //nolint:gosec
	if err != nil {
		return nil, fmt.Errorf("failed to read from %s: %w", fullPath, err)
	}

	return data, nil
}

// readFileFromAssets läser en fil från det inbäddade filsystemet och returnerar innehållet som en byte-slice.
func readFileFromAssets(name string) ([]byte, error) {
	data, err := Assets.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s from embedded assets: %w", name, err)
	}

	return data, nil
}

// GetFS returnerar antingen det inbäddade filsystemet eller ett filsystem som pekar på "static/" katalogen.
func GetFS() fs.FS {
	if directoryExists("static") {
		return os.DirFS("static")
	}

	return Assets
}
