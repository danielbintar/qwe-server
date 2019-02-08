class CreateCharacter < ActiveRecord::Migration[5.2]
  def change
    create_table :characters, id: :unsigned_integer do |t|
      t.integer :user_id, null: false, unsigned: true
      t.string :name, null: false

      t.timestamps
    end
  end
end
