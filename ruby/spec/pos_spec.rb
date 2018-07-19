require 'rspec'
require './pos'

describe Pos do
    describe '.barcode' do
        it 'returns a price when a matching barcode is found' do
            pos = Pos.new()
            pos.barcode('1234567890')

            expect(pos.message).to eq '$1.00'
        end
    end
end