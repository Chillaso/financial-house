db = new Mongo().getDB("financialHouse");
db.createCollection('entry', {capped: false});
db.entry.insertMany([
    {
        "year": 2020,
        "month": 11,
        "date": "2020-11-28T00:00:00Z",
        "type": 0,
        "amount": 50,
        "concept": "Facturas",
        "category": "Gasto fijo",
        "description": "Gas"
    },
    {
        "year": 2020,
        "month": 12,
        "date": "2020-11-28T00:00:00Z",
        "type": 0,
        "amount": 400,
        "concept": "Alquiler",
        "category": "Gasto fijo",
        "description": "Alquiler"
    }
]);
    /*{
        "year": 2020,
        "month": 12,
        "date": "2020-11-28T00:00:00Z",
        "type": 0,
        "amount": 400,
        "concept": "Alquiler",
        "category": "Gasto fijo",
        "description": "Alquiler"
    }
     */