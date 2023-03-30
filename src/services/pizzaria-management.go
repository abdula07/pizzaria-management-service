package services

import (
	"pizzeria-management-service/src/dbMain"
	"pizzeria-management-service/src/tracer"
)

type Order struct {
	Id    int64           `json:"id"`
	Done  bool            `json:"done"`
	Items []OrderProducts `json:"items"`
}

type OrderProducts struct {
	Id        int64 `json:"id"`
	OrderId   int64 `json:"orderId"`
	ProductId int64 `json:"productId"`
}

var Orders Order

func (obj *Order) Orders(products []OrderProducts) any {
	if len(products) <= 0 {
		return "items empty"
	}
	order := Order{Done: false}
	queryOrderInsert, errQueryOrderInsert := dbMain.Db.Exec("INSERT INTO `order` () VALUES ()")
	if errQueryOrderInsert != nil {
		tracer.Error(errQueryOrderInsert.Error())
	}
	orderId, errLastInsertId := queryOrderInsert.LastInsertId()
	if errLastInsertId != nil {
		tracer.Error(errLastInsertId.Error())
	}
	order.Id = orderId
	for key := range products {
		product := &products[key]
		product.OrderId = order.Id
		_, errQueryProductInsert := dbMain.Db.Query("INSERT INTO `order_products` (`order_id`, `product_id`) VALUES (?, ?);", order.Id, product.ProductId)
		if errQueryProductInsert != nil {
			tracer.Error(errQueryProductInsert.Error())
		}
	}
	order.Items = products
	return order
}

func (obj *Order) Products(products []OrderProducts, orderId string) any {
	if len(products) <= 0 {
		return "items empty"
	}
	queryOrderGet := dbMain.Db.QueryRow("SELECT `id`, `done` FROM `order` WHERE id=?", orderId)
	order := Order{}
	errScanOrder := queryOrderGet.Scan(&order.Id, &order.Done)
	if errScanOrder != nil {
		tracer.Error(errScanOrder.Error())
		return "Order not found"
	}
	if order.Done == true {
		return "Order already done"
	}
	for key := range products {
		product := products[key]
		product.OrderId = order.Id
		_, errQueryProductInsert := dbMain.Db.Query("INSERT INTO `order_products` (`order_id`, `product_id`) VALUES (?, ?);", order.Id, product.ProductId)
		if errQueryProductInsert != nil {
			tracer.Error(errQueryProductInsert.Error())
		}
	}
	order.Items = products
	return true
}

func (obj *Order) Get(orderId string) any {
	queryOrderGet := dbMain.Db.QueryRow("SELECT `id` FROM `order` WHERE id=?", orderId)
	order := Order{}
	errScanOrder := queryOrderGet.Scan(&order.Id)
	if errScanOrder != nil {
		tracer.Error(errScanOrder.Error())
		return "Order not found"
	}
	orderProducts, errOrderProducts := dbMain.Db.Query("SELECT `id`, `order_id`, `product_id` FROM `order_products` WHERE order_id=?", orderId)
	if errOrderProducts != nil {
		tracer.Error(errOrderProducts.Error())
	}
	for orderProducts.Next() {
		var orderProduct OrderProducts
		errScanOrderProduct := orderProducts.Scan(&orderProduct.Id, &orderProduct.OrderId, &orderProduct.ProductId)
		if errScanOrderProduct != nil {
			tracer.Error(errScanOrderProduct.Error())
		}
		order.Items = append(order.Items, orderProduct)
	}
	return order
}

func (obj *Order) DoneMake(orderId string) any {
	queryOrderGet := dbMain.Db.QueryRow("SELECT `id`, `done` FROM `order` WHERE id=?", orderId)
	order := Order{}
	errScanOrder := queryOrderGet.Scan(&order.Id, &order.Done)
	if errScanOrder != nil {
		tracer.Error(errScanOrder.Error())
		return "Order not found"
	}
	if order.Done == true {
		return "Order already done"
	}
	_, errQueryOrderUpdate := dbMain.Db.Query("UPDATE `order` SET `done` = true WHERE `id` = ?;", order.Id)
	if errQueryOrderUpdate != nil {
		tracer.Error(errQueryOrderUpdate.Error())
		return "Error update"
	}
	return true
}

func (obj *Order) All(done bool, isCondition bool) []Order {
	queryOrdersGet, errQueryOrdersGet := dbMain.Db.Query("SELECT `id`, `done` FROM `order`")
	if isCondition == true {
		queryOrdersGet, errQueryOrdersGet = dbMain.Db.Query("SELECT `id`, `done` FROM `order` WHERE done=?", done)
	}
	if errQueryOrdersGet != nil {
		tracer.Error(errQueryOrdersGet.Error())
	}
	var orders []Order
	for queryOrdersGet.Next() {
		order := Order{}
		errScanOrder := queryOrdersGet.Scan(&order.Id, &order.Done)
		if errScanOrder != nil {
			tracer.Error(errScanOrder.Error())
		}
		orders = append(orders, order)
	}
	return orders
}
