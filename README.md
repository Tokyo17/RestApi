# Projct Api Ecommerce Sederhana
Project ini adalah api tentang penjualan pakaian seperti kaos dan sweater di api ini memeliki beberapa fitur seperti regsiter,login,upload product, transaksi dan sebaginya.


# MVP #

* Register, login dan menampilkan informasi user
* Menampilkan Produk
* Pencarian produk berdasarkan type
* Melakukan transaksi
* Pembayaran
* Pengiriman dan menampilkan informasinya
* Menampilkan informasi tentang produk yang dibeli sebelumnya
* Upload dan update produk



# Daftar Api #

## admin ##
* admins/registers : register admin
* admins/logins : login admin

## user ##
* users/logins : login user
* users/details/:id : mendapatkan detail user berdasarkan id
* users/registers : register user



## product ##
* products : menampilkan semua produk
* products/:id : memfilter produk berdasarkan product type id
* products : mengupload produk baru (admin)
* products : mengupdate produk yang sudah ada (admin)
* products/types : mengupload type produk baru (admin)

## transactions ##
* transactions : menambahkan produk ke keranjang
* transactions/details : menampilkan detail keranjang belanja
* transactions/pay : melakukan pembayaran sesuai dengan nominal yang ada di transactions
* payments : menambahkan pilihan pembayaran (admin)
* payments : menampilkan semua pilihan pembayaran
* shipments : menambahakan pilihan pengiriman (admin)
* shipments : menampilkan semua pilihan pengiriman
* transactions/transactiondetails : menampilkan detail transaksi setelah melakukan pembayaran


