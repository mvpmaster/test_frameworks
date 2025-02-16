
let test = 'test'.repeat(5000)

Bun.serve({
    fetch(request){
        return new Response(test)
    }
})
console.log("Listening on Port 3000")