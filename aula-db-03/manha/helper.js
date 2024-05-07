let mongoQuery;
const db = () => {
  return {
    restaurants: {
      findOne: () => { },
      find: () => { },
      aggregate: () => { },
    }
  }
}



mongoQuery = db.restaurants.findOne({}, {
  restaurante_id: 1,
  nombre: 1,
  barrio: 1,
  tipo_cocina: 1,
  _id: 0,
})

mongoQuery = db.restaurants.find(
  { nombre: { $regex: /Bake/ } },
  {
    restaurante_id: 1,
    barrio: 1,
    tipo_cocina: 1,
    _id: 0,
  }
).limit(3)

mongoQuery = db.restaurants.find(
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

mongoQuery = db.restaurants.find(
  {
    tipo_cocina: {
      $or: [
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

mongoQuery = db.restaurants.find(
  {
    grados: {
      $elemMatch: {
        grado: "A",
        puntaje: { $gt: 20 }
      }
    }
  }
).limit(3)

mongoQuery = db.restaurants.find(
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

mongoQuery = db.restaurants.find(
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

mongoQuery = db.restaurants.aggregate([
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

mongoQuery = db.restaurants.aggregate([
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

mongoQuery = db.restaurants.aggregate([
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

mongoQuery = db.restaurants.find(
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