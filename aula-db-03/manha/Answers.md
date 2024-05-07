# Bancos de dados não-relacionais

## Aula 03 - Manhã

### Configurações

- Iniciei uma database local

- CONN_STRING: mongodb://localhost:27017/bootcamp_exercises

- Importei os dados da collection.json para a collection restaurants da database exercises

### Para começar

1. Quantas coleções tem o banco de dados?

- Dentro da minha database "exercise" apenas uma collection "restaurants"

2. Quantos documentos em cada coleção? Quanto ocupa cada coleção?

- 25359 documentos, 11140976 bytes

3. Quantos índices em cada coleção? Quanto espaço os índices de cada coleção ocupam?

- Apenas o _id
- 266240 bytes total, dividido pelo total de indexes temos... o mesmo número

4. Traga um documento de exemplo de cada coleção. db.collection.find(...).pretty() nos dá um formato mais legível.

```js
  {
    _id: ObjectId('5eb3d668b31de5d588f42a75'),
    direccion: {
      edificio: '2055',
      coord: [ -74.1321, 40.61266000000001 ],
      calle: 'Victory Boulevard',
      codigo_postal: '10314'
    },
    barrio: 'Staten Island',
    tipo_cocina: 'American',
    grados: [
      {
        date: ISODate('2014-11-06T00:00:00.000Z'),
        grado: 'B',
        puntaje: 25
      },
      {
        date: ISODate('2014-05-06T00:00:00.000Z'),
        grado: 'B',
        puntaje: 20
      },
      {
        date: ISODate('2013-01-26T00:00:00.000Z'),
        grado: 'A',
        puntaje: 13
      },
      {
        date: ISODate('2011-12-17T00:00:00.000Z'),
        grado: 'A',
        puntaje: 7
      }
    ],
    nombre: "Schaffer'S Tavern",
    restaurante_id: '40371771'
  }
  ```

5. Para cada coleção, liste os campos de nível raiz (ignore campos em documentos aninhados) e seus tipos de dados.

- Ignorando os campos aninhados:

```txt
  _id: ObjectId
  direccion: Object
  barrio: String
  tipo_cocina: String
  grados: Array
  nombre: String
  restaurante_id: String
```

### Exercício 1: SQL

1. Devolver restaurante_id, nombre, barrio e tipo_cocina mas excluindo _id para um documento (o primeiro).

Query:

```js
db.restaurants.findOne({}, {
  restaurante_id: 1,
  nombre: 1,
  barrio: 1,
  tipo_cocina: 1,
  _id: 0,
})
```

Result:

```js
{
  barrio: 'Manhattan',
  tipo_cocina: 'American',
  nombre: 'Cafe Metro',
  restaurante_id: '40363298'
}
```

2. Devolver restaurante_id, nombre, barrio e tipo_cocina para os primeiros 3 restaurantes que contenham 'Bake' em alguma parte do seu nome.

```js
db.restaurants.find(
  { nombre: { $regex: /Bake/ } },
  {
    restaurante_id: 1,
    barrio: 1,
    tipo_cocina: 1,
    _id: 0,
  }
).limit(3)
```

Result:

```js
[
  {
    barrio: 'Staten Island',
    tipo_cocina: 'American',
    restaurante_id: '40370910'
  },
  {
    barrio: 'Queens',
    tipo_cocina: 'Caribbean',
    restaurante_id: '40377560'
  },
  {
    barrio: 'Bronx',
    tipo_cocina: 'Bakery',
    restaurante_id: '30075445'
  }
]
```

3. Contar os restaurantes de comida (tipo_cocina) china (Chinese) o tailandesa (Thai) do bairro (bairro) Bronx. Consultar or versus in.

Query

```js
db.restaurants.find(
  {
    tipo_cocina: {
      $in: [
        "Chinese",
        "Thai"
      ]
    },
    barrio: {
      $in: [
        "Bronx",
      ]
    }
  },
).count()
```

Results: 325


### Exercício 2: NoSQL


1. Traga 3 restaurantes que receberam pelo menos uma classificação de grau 'A' com pontuação maior que 20. A mesma classificação deve atender às duas condições simultaneamente; investigue o operador elemMatch.

Query:

```js
db.restaurants.find(
  {
    grados: {
      $elemMatch: {
        grado: "A",
        puntaje: { $gt: 20 }
      }
    }
  }
).limit(3)
```

Result (with limit 1 to simplify):

```js
[
  {
    _id: ObjectId('5eb3d668b31de5d588f42eec'),
    direccion: {
      edificio: '107-23',
      coord: [ -73.834012, 40.683833 ],
      calle: 'Liberty Avenue',
      codigo_postal: '11417'
    },
    barrio: 'Queens',
    tipo_cocina: 'Caribbean',
    grados: [
      {
        date: ISODate('2014-03-29T00:00:00.000Z'),
        grado: 'A',
        puntaje: 27
      },
      {
        date: ISODate('2013-06-12T00:00:00.000Z'),
        grado: 'A',
        puntaje: 12
      },
      {
        date: ISODate('2012-05-10T00:00:00.000Z'),
        grado: 'A',
        puntaje: 13
      },
      {
        date: ISODate('2011-12-29T00:00:00.000Z'),
        grado: 'A',
        puntaje: 13
      }
    ],
    nombre: "Gemini'S Lounge",
    restaurante_id: '40511696'
  }
]
```

2. Quantos documentos estão faltando coordenadas geográficas? Em outras palavras, verifique se o tamanho de direccion.coord é 0 e contar.

R: 2

Query:

```js
db.restaurants.find(
  {
    $or: [{
      "direccion.coord": {
        $size: 0
      }
    }, {
      "direccion.coord": {
        $exists: false
      }
    }]
  }
).count()
```

3. Devolver nombre, barrio, tipo_cocina y grados para os 3 primeiros restaurantes; de cada documento apenas a última avaliação. Ver o operador slice.

Query:

```js
db.restaurants.find(
  {},
  {
    nombre: 1,
    barrio: 1,
    tipo_cocina: 1,
    grados: {
      $slice: -1
    },
    _id: 0,
  }
).limit(3)
```

Mas, particularmente acho que ao invés de retornar o ultimo grado, faz mais sentido retornar baseado em grados ordenado pela data, retornando o mais recente:

```js
db.restaurants.aggregate([
  {
    $unwind: "$grados"
  },
  {
    $sort: { "grados.index.date": -1 }
  },
  {
    $group: {
      _id: "$_id",
      nombre: { $first: "$nombre" },
      barrio: { $first: "$barrio" },
      tipo_cocina: { $first: "$tipo_cocina" },
      grados: { $first: "$grados" },
    },
  },
  {
    $project: {
      _id: 0,
      nombre: 1,
      barrio: 1,
      tipo_cocina: 1,
      grados: 1,
    }
  },
  {
    $limit: 3
  }
])
```

### Exercício 3: Potpourri

1. Quais são os 3 principais tipos de culinária (cuisine) que podemos encontrar nos dados? Googlear "mongodb group by field, count it and sort it". Ver a etapa limit do pipeline de agregação.

R: 

Query:

```js
db.restaurants.aggregate([
  {
    $group: {
      _id: "$tipo_cocina",
      count: { $sum: 1 },
    },
  },
  {
    $sort: { count: -1 }
  },
  {
    $limit: 3
  }
])
```

Result:
```js
[
  { _id: 'American', count: 6183 },
  { _id: 'Chinese', count: 2418 },
  { _id: 'Café/Coffee/Tea', count: 1214 }
]
```

2. Quais são os bairros mais desenvolvidos gastronomicamente? Calcular a média ($avg) a pontuação (grades.score) por bairro; considerando restaurantes com mais de três avaliações; oClassifique os bairros com a melhor pontuação. Ajuda:

2.a match é um estágio que filtra documentos com base em uma condição, similar a db.orders.find (<condição>).
2.b Parece necessário desconstruir as listas grades para produzir um documento para cada ponto usando o palco unwind.

R: 

Query:
```js
db.restaurants.aggregate([
  {
    $unwind: "$grados"
  },
  {
    $group: {
      _id: "$barrio",
      count: { $sum: 1 },
      avgPuntaje: { $avg: "$grados.puntaje" },
    },
  },
  {
    $sort: { avgPuntaje: -1 }
  },
  {
    $limit: 3
  }
])
```

Result:
```js
[
  { _id: 'Queens', count: 20877, avgPuntaje: 11.634865110930088 },
  { _id: 'Brooklyn', count: 21963, avgPuntaje: 11.44797595737899 },
  { _id: 'Manhattan', count: 38622, avgPuntaje: 11.418151216986018 }
]
```

3. Uma pessoa querendo comer está na longitude -73.93414657 e na latitude 40.82302903, Quais opções você tem em um raio de 500 metros? Consultar geospatial tutorial.

R:

Tive que criar um index:

```js
db.restaurants.createIndex({ "direccion.coord" : "2dsphere" })
```

Query:

```js
db.restaurants.find(
  {
    "direccion.coord": {
      $nearSphere: {
        $geometry: {
          type: "Point",
          coordinates: [-73.93414657, 40.82302903]
        },
        $maxDistance: 500
      }
    }
  }
).count()
```

Result: 24