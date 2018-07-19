class Pos
    attr_reader :message

    def initialize
        @message = ''
        @product_store = { '1234567890' => '$1.00' }
    end

    def barcode(barcode_input)
        @message = @product_store[barcode_input]
    end
end