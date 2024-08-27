const express = require('express');
const app =  express();
const port = 3000

const config = {
    host: 'db',
    user: 'root',
    password: 'root',
    database:'nodedb',
    port: 3306
};

const mysql = require('mysql')
const connection = mysql.createConnection(config)

const sql = `INSERT INTO people(name) values('Wesley')`
connection.query(sql)
connection.end()
console.log("Inserido com sucesso");


app.get('/healthcheck',(req,res)=>{
    res.send('<h1>Full Cycle Deu certo a parada aqui <h1>')
});
console.log("Teste");

app.listen(port,()=>{
    console.log(`Rodando na porta ${port}`);
})
