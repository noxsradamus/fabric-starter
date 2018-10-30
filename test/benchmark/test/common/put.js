'use strict';

module.exports.info  = 'put';

let bc, contx, timeout;
module.exports.init = function(blockchain, context, args = {}) {
    bc = blockchain;
    contx = context;
    timeout = args.timeout || 60;
    return Promise.resolve();
};

const ids = [];
module.exports.ids = ids;

function _generateId() {
    return `${Math.random().toString().substring(2, 8)}-${Date.now()}`;
}

module.exports.run = function() {
    const nextId = _generateId();
    return bc.invokeSmartContract(contx, 'reference', 'v1',
        [{transaction_type: 'put', id: nextId, value: "test"}],
        timeout)
      .then(res => {
          if(res[0] && res[0].status && res[0].status.status === 'success') {
              ids.push(nextId);
          }
          return res;
      });
};

module.exports.end = function() {
    return Promise.resolve();
};
