let lastMessage = ''

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
