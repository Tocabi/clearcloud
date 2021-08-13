// https://developer.mozilla.org/en-US/docs/Web/API/FileSystemEntry
export interface FileSystemEntry {
  isFile: boolean;
  isDirectory: boolean;
  name: string;
  fullPath: string;
}

export interface FileSystemDirectoryEntry extends FileSystemEntry {
  createReader: () => FileSystemDirectoryReader;
}

export interface FileSystemFileEntry extends FileSystemEntry {
  file: (success: (file: File) => void, error: (error: unknown) => void) => void;
}

export interface FileSystemDirectoryReader {
  readEntries: (success: (entries: FileSystemEntry[]) => void, error: (error: unknown) => void) => void;
}

export function isFileEntry(entry: FileSystemEntry): entry is FileSystemFileEntry {
  return entry.isFile;
}

export function isDirectoryEntry(entry: FileSystemEntry): entry is FileSystemDirectoryEntry {
  return entry.isDirectory;
}

export function getFsEntryFile(entry: FileSystemFileEntry): Promise<File> {
  return new Promise((resolve, reject) => entry.file(resolve, reject));
}

export async function getAllEntriesFromTree(items: DataTransferItemList): Promise<FileSystemEntry[]> {
  const result: FileSystemEntry[] = [];

  for (const item of Array.from(items)) {
    const entry: FileSystemEntry = item.webkitGetAsEntry();
    if (!entry) {
      continue;
    }

    result.push(entry);
    if (isDirectoryEntry(entry)) {
      result.push(...(await readAllEntriesRecursive(entry)));
    }
  }

  return result;
}

async function readAllEntriesRecursive(dir: FileSystemDirectoryEntry): Promise<FileSystemEntry[]> {
  const result = await readAllEntries(dir);

  for (const entry of result) {
    if (isDirectoryEntry(entry)) {
      result.push(...(await readAllEntriesRecursive(entry)));
    }
  }

  return result;
}

async function readAllEntries(dir: FileSystemDirectoryEntry): Promise<FileSystemEntry[]> {
  const result: FileSystemEntry[] = [];
  const reader = dir.createReader();

  let read = await readEntries(reader);
  while (read.length) {
    result.push(...read);
    read = await readEntries(reader);
  }

  return result;
}

function readEntries(reader: FileSystemDirectoryReader): Promise<FileSystemEntry[]> {
  return new Promise((resolve, reject) => reader.readEntries(resolve, reject));
}
