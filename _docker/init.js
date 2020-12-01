db = new Mongo().getDB("financialHouse");
db.createCollection('book', {capped: false});
db.book.insert({
  "year": 2020,
  "months": [{
    "month": 11,
    "entries": [{
      "date": "2020-11-28T00:00:00Z",
      "type": 0,
      "amount": 50,
      "concept": "Gasto fijo",
      "description": "Gas"
    }]
  }]
})