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
console.log('Tentando conectar');
const connection = mysql.createConnection(config)
const sqlCreateTable = `CREATE TABLE IF NOT EXISTS people(id int not null auto_increment, name varchar(255), primary key(id))`

connection.query(sqlCreateTable);
console.log('Tabela criada');
connection.end()

app.get('/',(req,res)=>{
    const connection = mysql.createConnection(config)
    const sqlGet = `SELECT id,name FROM people`;
    let count = 0;
    let data;
    console.log('Entrou');
    connection.query(sqlGet,(err,rows,fields)=>{
        if(err) throw err;
        if(!rows){
            data = 'Nenhum registro encontrado';
        }
        let names = '';
        rows.forEach(row => {
            names += row.id + ' - ' + row.name + '<br>';
        });
        data+="\n" +names;
    });
    console.log(data);
    const sql = `INSERT INTO people(name) values('${req.params.name==undefined?'Teste 1':req.params.name}')`
    connection.query(sql);
    connection.end()

    res.send('<h1>Full Cycle Rocks!<h1>\n'+data);
});


app.listen(port,()=>{
    console.log(`Rodando na porta ${port}`);
    console.log(app.locals);
})
