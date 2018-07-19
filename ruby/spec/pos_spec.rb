require 'rspec'
require './pos'

describe Pos do
    describe '.barcode' do
        it 'returns a price when a matching barcode is found' do
            pos = Pos.new()
            pos.barcode("1234567890\n\r")

            expect(pos.message).to eq '$1.00'
        end

        it 'returns an error message when a barcode isn\'t found' do
            pos = Pos.new()
            pos.barcode("alpha\n\r")

            expect(pos.message).to eq 'Unable to locate price for barcode'
        end
    end
end