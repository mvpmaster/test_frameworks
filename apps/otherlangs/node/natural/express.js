
const express = require('express')
const app = express()
const port = 3000

test = 'test'.repeat(5000)

app.get('/', (req, res) => {
  res.send(test) //'Hello World!')
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})