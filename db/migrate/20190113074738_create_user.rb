class CreateUser < ActiveRecord::Migration[5.2]
  def change
    create_table :users, id: :unsigned_integer do |t|
      t.string :username, null: false
      t.string :password, null: false

      t.timestamps
    end
  end
end
