const cero = require('0http')
const { router, server } = cero()

test = 'test'.repeat(5000)
test2 = '{"error":"'+test+'"}'

router.get('/test', (req, res) => {
  res.end(test2)
})

// router.post('/test', (req, res) => {
//   // ...
//   res.statusCode = 200
//   res.end("helo")
// })

//...

server.listen(3000)