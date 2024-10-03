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

function rodarScript(){
    
    const connection = mysql.createConnection(config)
    const sqlCreateTable = `CREATE TABLE IF NOT EXISTS people(id int not null auto_increment, name varchar(255), primary key(id))`    
    connection.query(sqlCreateTable);
    console.log('Tabela criada');
    connection.commit();
}

app.get('/',(req,res)=>{   
    
    rodarScript(); 
    const connection = mysql.createConnection(config)
    const sqlGet = `SELECT id,name FROM people`;    
    let html = '<h1>Full Cycle Rocks!</h1></br>';    
    
    connection.query(sqlGet, (error, results, fields) => {
        if (error) {
            console.error('Error executing query:', error);
            return;
        }
        results.forEach(row => {            
            html += row['id'] + ' - ' + row['name'] + '</br>';            
        });

        res.send(html);
    });
    
    const sql = `INSERT INTO people(name) values('Teste 1')`;
    connection.query(sql);
    connection.commit();
    connection.end();

});


app.listen(port,()=>{
    console.log(`Rodando na porta ${port}`);
    console.log(app.locals);
})
