const express = require('express');
const FileStore = require('./storage');
const routes = require('./routes');

const app = express();
const fileStore = new FileStore();

app.use(express.json());
app.use(express.raw({ type: 'application/octet-stream', limit: '10mb' }));

routes(app, fileStore);

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
  console.log(`File store server running on port ${PORT}`);
});

module.exports = { app, fileStore };