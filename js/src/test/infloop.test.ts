import * as assert from 'assert';
import * as test from '../test';

const Test = test.Test();

describe.skip('Really Long Loop', function () {
  this.timeout(10 * 1000)
  let contract

  before(Test.before(async function (burrow) {
    const source = `
      pragma solidity >=0.0.0;
      contract main {
        function test() public returns (string memory) {
            c sub = new c();
            return sub.getString();
          }
        }
        contract c {
        string s = "secret";
        uint n = 0;
        function getString() public returns (string memory){
          for (uint i = 0; i < 10000000000000; i++) {
            n += 1;
          }
          return s;
        }
      }
    `

    const {abi, bytecode} = test.compile(source, 'main')
    const c = await burrow.contracts.deploy(abi, bytecode);
    contract = c;
  }))

  after(Test.after())

  it('It catches a revert when gas runs out',
    Test.it(function () {
      return contract.test()
        .then((str) => {
          throw new Error('Did not catch revert error')
        }).catch((err) => {
          assert.equal(err.message, 'ERR_EXECUTION_REVERT')
        })
    })
  )
})
