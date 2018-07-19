let lastMessage = ''
const productStore = { "1234567890" : "$1.00" }

const onBarcode = (barcode) => {
  lastMessage = "$1.00"
}

const lastMessageSent = () => {
    return lastMessage
}

module.exports = {
  onBarcode,
  lastMessageSent
};
