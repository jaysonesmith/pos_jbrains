const poc = require("../pos");
const expect = require("chai").expect;

describe("pos", () => {
  it("returns a price if a matching barcode is found", () => {
    pos.onBarcode("1234567890");
    expect(pos.lastMessage).to.equal("$1.00")
  });
});
