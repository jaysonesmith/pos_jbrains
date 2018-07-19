const {onBarcode, lastMessageSent} = require("../src/pos");
const expect = require("chai").expect;

describe("pos", () => {
  it("returns a price if a matching barcode is found", () => {
    onBarcode("1234567890");
    expect(lastMessageSent()).to.equal("$1.00")
  });
});
