const express = require('express')
const app =  express();
const port = 3000

app.get('/',(req,res)=>{
    res.send('<h1>Full Cycle<h1>')
});
console.log("Teste");

app.listen(port,()=>{
    console.log(`Rodando na porta ${port}`);
})
