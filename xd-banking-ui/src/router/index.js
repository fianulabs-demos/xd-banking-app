import { route } from 'quasar/wrappers'
import { createRouter, createMemoryHistory, createWebHistory, createWebHashHistory } from 'vue-router'
import routes from './routes'

/*
 * If not building with SSR mode, you can
 * directly export the Router instantiation;
 *
 * The function below can be async too; either use
 * async/await or return a Promise which resolves
 * with the Router instance.
 */

export default route(function (/* { store, ssrContext } */) {
  const createHistory = process.env.SERVER
    ? createMemoryHistory
    : (process.env.VUE_ROUTER_MODE === 'history' ? createWebHistory : createWebHashHistory)

  const Router = createRouter({
    scrollBehavior: () => ({ left: 0, top: 0 }),
    routes,

    // Leave this as is and make changes in quasar.conf.js instead!
    // quasar.conf.js -> build -> vueRouterMode
    // quasar.conf.js -> build -> publicPath
    history: createHistory(process.env.MODE === 'ssr' ? void 0 : process.env.VUE_ROUTER_BASE)
  })

  return Router
})


const express = require('express');
const app = express();
const mysql = require('mysql');

// Set up the MySQL connection
const connection = mysql.createConnection({
  host: 'localhost',
  user: process.env.USER,
  password: process.env.PASSWORD,
  database: process.env.DATABASE
});

// Connect to the database
connection.connect();

// A route that takes a user ID as a parameter and returns user details
app.get('/user/:id', (req, res) => {
  // Vulnerable code: Directly concatenating user input into the SQL query
  const userId = req.params.id;
  const query = `SELECT * FROM users WHERE id = '${userId}'`;

  connection.query(query, (error, results) => {
    if (error) {
      res.status(500).send('Server error');
      return;
    }
    res.send(results);
  });
});

// Start the server
app.listen(3000, () => {
  console.log('Server running on port 3000');
});
