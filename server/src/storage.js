const crypto = require('crypto');

class FileStore {
  constructor() {
    this.files = new Map();
  }

  calculateChecksum(content) {
    return crypto.createHash('sha256').update(content).digest('hex');
  }

  addFile(filename, content) {
    const checksum = this.calculateChecksum(content);
    
    for (const [existingFile, existingContent] of this.files) {
      const existingChecksum = this.calculateChecksum(existingContent);
      if (existingChecksum === checksum) {
        return { success: true, message: 'File content already exists', existingFile };
      }
    }

    if (this.files.has(filename)) {
      return { success: false, message: 'File already exists' };
    }

    this.files.set(filename, content);
    return { success: true, message: 'File added successfully' };
  }

  listFiles() {
    return Array.from(this.files.keys());
  }

  removeFile(filename) {
    if (!this.files.has(filename)) {
      return { success: false, message: 'File not found' };
    }
    this.files.delete(filename);
    return { success: true, message: 'File removed successfully' };
  }

  updateFile(filename, content) {
    this.files.set(filename, content);
    return { success: true, message: 'File updated successfully' };
  }

  getWordCount() {
    let totalWords = 0;
    for (const content of this.files.values()) {
      totalWords += content.split(/\s+/).filter(word => word.length > 0).length;
    }
    return totalWords;
  }

  getFrequentWords(limit = 10, order = 'dsc') {
    const allWords = [];
    for (const content of this.files.values()) {
      allWords.push(...content.toLowerCase().split(/\s+/).filter(word => word.length > 0));
    }

    const wordFreq = allWords.reduce((acc, word) => {
      acc[word] = (acc[word] || 0) + 1;
      return acc;
    }, {});

    const sortedWords = Object.entries(wordFreq)
      .sort((a, b) => order === 'dsc' ? b[1] - a[1] : a[1] - b[1])
      .slice(0, limit);

    return sortedWords;
  }
}

module.exports = FileStore;