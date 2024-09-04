package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Produk dengan nama dan harga
type Product struct {
	Name  string
	Price float64
}

// Keranjang belanja
type Cart struct {
	Items []Product
}

// Menambahkan produk ke dalam keranjang
func (c *Cart) AddProduct(p Product) {
	c.Items = append(c.Items, p)
}

// Menghitung total harga
func (c *Cart) TotalPrice() float64 {
	total := 0.0
	for _, item := range c.Items {
		total += item.Price
	}
	return total
}

// Menampilkan struk belanja
func (c *Cart) PrintReceipt() {
	fmt.Println("\n--- STRUK BELANJA ---")
	for _, item := range c.Items {
		fmt.Printf("%s: $%.2f\n", item.Name, item.Price)
	}
	fmt.Printf("TOTAL: $%.2f\n", c.TotalPrice())
	fmt.Println("----------------------")
}

func main() {
	// Daftar produk yang tersedia
	products := []Product{
		{"Apple", 20000},
		{"Banana", 12000},
		{"Orange", 17000},
		{"Milk", 18000},
		{"Bread", 15000},
	}

	cart := Cart{}

	for {
		fmt.Println("\nSelamat datang di POS!")
		fmt.Println("Daftar produk yang tersedia:")
		for i, p := range products {
			fmt.Printf("%d. %s - Rp.%.2f\n", i+1, p.Name, p.Price)
		}

		fmt.Print("\nPilih produk (masukkan nomor produk) atau ketik 'checkout' untuk menyelesaikan: ")
		var input string
		fmt.Scanln(&input)

		if strings.ToLower(input) == "checkout" {
			cart.PrintReceipt()
			break
		}

		productIndex, err := strconv.Atoi(input)
		if err != nil || productIndex < 1 || productIndex > len(products) {
			fmt.Println("Nomor produk tidak valid, coba lagi.")
			continue
		}

		selectedProduct := products[productIndex-1]
		cart.AddProduct(selectedProduct)
		fmt.Printf("%s telah ditambahkan ke keranjang.\n", selectedProduct.Name)
	}
}
