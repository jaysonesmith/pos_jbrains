class Pos
    attr_reader :message

    def initialize
        @message = ''
    end

    def barcode(barcode_input)
        @message = '$1.00'
    end
end