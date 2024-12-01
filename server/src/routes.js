module.exports = (app, fileStore) => {
    // Add file
    app.post('/files', (req, res) => {
      const filename = req.headers['x-filename'];
      const content = req.body;
      const result = fileStore.addFile(filename, content.toString());
      console.log(`File is Added.`);
      res.json(result);
    });
  
    // List files
    app.get('/files', (req, res) => {
      res.json(fileStore.listFiles());
      console.log(`List of the Files:`);
    });
  
    // Remove file
    app.delete('/files/:filename', (req, res) => {
      const result = fileStore.removeFile(req.params.filename);
      console.log(`Files are Removed.`);
      res.json(result);
    });
  
    // Update file
    app.put('/files/:filename', (req, res) => {
      const filename = req.params.filename;
      const content = req.body;
      const result = fileStore.updateFile(filename, content.toString());
      console.log(`Files are Updated.`);
      res.json(result);
    });
  
    // Word count
    app.get('/files/wordcount', (req, res) => {
      res.json({ wordCount: fileStore.getWordCount() });
      console.log(`Word Count`);
    });
  
    // Frequent words
    app.get('/files/frequent-words', (req, res) => {
      const limit = parseInt(req.query.limit) || 10;
      const order = req.query.order || 'dsc';
      res.json(fileStore.getFrequentWords(limit, order));
      console.log(`Frequent Words :`);
    });
  };