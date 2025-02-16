const server = require('low-http-server')({})

test = 'test'.repeat(5000)

server.on('request', (req, res) => {
  res.end(test) //'Hello World!')
})



server.listen(3000, () => {
  console.log('Server listening on http://0.0.0.0:3000')
})
