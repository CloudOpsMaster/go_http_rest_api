import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Product } from './product-page/interfaces'

@Injectable({
  providedIn: 'root'
})
export class ProductService {

  apiUrl = 'http://127.0.0.1:8080/products';

  constructor(private http: HttpClient) {} 

  getProducts() {
    return this.http.get<Product[]>(this.apiUrl)
  }
}
