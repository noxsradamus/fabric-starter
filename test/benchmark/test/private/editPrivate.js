'use strict';

module.exports.info  = 'edit Private';

let bc, contx, ids, timeout;
module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;
    ids = require('./putPrivate').ids;
    timeout = args.timeout || 60;
    return Promise.resolve();
};

module.exports.run = function() {
    const nextId = ids[Date.now() % ids.length];
    return bc.invokeSmartContract(contx, 'reference', 'v1',
      [{transaction_type: 'editPrivate', id: nextId, value: "test_new"}],
      timeout);
};

module.exports.end = function() {
    return Promise.resolve();
};
