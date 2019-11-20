var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  setTimeout(() => {
    console.log("TIMER COMES OUT!");
  }, 3000);
  res.render('index', { title: 'Express' });
});

module.exports = router;
