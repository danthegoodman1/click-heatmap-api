const duckdb = require('duckdb')

async function main() {
  const db = new duckdb.Database(':memory:')
  db.all(`INSTALL 'json'`, function(err: Error, res: any) {
    if (err) {
      throw err;
    }
    console.log(res)
  })
  db.all(`LOAD 'json'`, function(err: Error, res: any) {
    if (err) {
      throw err;
    }
    console.log(res)
  })
  db.all(`select * from duckdb_extensions()`, function(err: Error, res: any) {
    if (err) {
      throw err;
    }
    console.log(res)
  })
  db.all('SELECT 42 AS fortytwo', function(err: Error, res: any) {
    if (err) {
      throw err;
    }
    console.log(res[0].fortytwo)
  })
}

main()
