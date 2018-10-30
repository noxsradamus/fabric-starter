'use strict';

module.exports.info  = 'query by ID Private (CouchDB)';

let bc, contx, ids, timeout;
module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;
    ids = require('../putPrivate').ids;
    timeout = args.timeout || 60;
    return Promise.resolve();
};

module.exports.run = function() {
    const nextId = ids[Date.now() % ids.length];
    return bc.queryState(contx, 'reference', 'v1', [nextId], 'queryCouchPrivate', timeout)
};

module.exports.end = function() {
    return Promise.resolve();
};
