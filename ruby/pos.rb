class Pos
    attr_reader :message

    def initialize
        @message = ''
        @product_store = { '1234567890' => '$1.00' }
    end

    def barcode(barcode_input)
        price = @product_store.dig(barcode_input.strip!)
        
        @message = price || 'Unable to locate price for barcode'

        outputMessage(@message)
    end

    def outputMessage(message)
        puts message
    end
end