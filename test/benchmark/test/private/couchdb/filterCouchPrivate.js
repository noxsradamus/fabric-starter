'use strict';

module.exports.info  = 'filter Private (CouchDB)';


let bc, contx, timeout;
module.exports.init = function(blockchain, context, args) {
    bc       = blockchain;
    contx    = context;
    timeout = args.timeout || 60;
    return Promise.resolve();
};

module.exports.run = function() {
    return bc.queryState(contx, 'reference', 'v1', [], 'filterCouchPrivate', timeout);
};

module.exports.end = function() {
    // do nothing
    return Promise.resolve();
};
