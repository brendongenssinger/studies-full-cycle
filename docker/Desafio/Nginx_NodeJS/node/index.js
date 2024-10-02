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


app.get('/healthcheck',(req,res)=>{
    res.send('<h1>Full Cycle Deu certo a parada aqui <h1>')
});

app.post('/create',(req,res)=>{
    const sql = `INSERT INTO people(name) values('${req.body.name}')`
    connection.query(sql)  
    console.log("Inserido com sucesso");
});


app.listen(port,()=>{
    console.log(`Rodando na porta ${port}`);
    console.log(app.locals);
})
