package items

import (
	"bookstore/items-api/datasources"
	"bookstore/items-api/utils/errors"
	"fmt"
	"strings"
)

const (
	queryInsertItem         = "INSERT INTO items(title, seller, price, stock, sold_quantity, createdAt) VALUES(?, ?, ?, ?, ?, ?);"
	queryUpdateItem 		= "UPDATE items SET title=?, seller=?, price=?, stock=?, sold_quantity=? WHERE id = ?;"
	queryById 				= "SELECT title, seller, price, stock, sold_quantity, createdAt from items where id=?;"
	queryDeleteItem 		= "DELETE FROM items WHERE id = ?;"
	queryListItems			= "SELECT id, title, seller, price, stock, sold_quantity FROM items order by createdAt desc;"
	queryGetItemsByUserId   = "SELECT id, title, seller, price, stock, sold_quantity FROM items WHERE seller = ? order by createdAt desc;"
	// queryListProducts 			= "SELECT * from products;"
	// queryDeleteProduct = "DELETE FROM products WHERE id=?;"
	
)

func (item *Item) GetByUserId() ([]*Item, *errors.RestErr) {
	stmt, err := datasources.Client.Prepare(queryGetItemsByUserId)
	if err != nil {
		return nil, errors.NewInternalServerError("error when tying to get item" + err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(item.Seller) //update with limit and skip
	if err != nil {
		return nil, errors.NewInternalServerError("error when tying to get cashier" + err.Error())
	}
	defer rows.Close()

	results := make([]*Item, 0)
	for rows.Next() {
		x := &Item{}
		if err := rows.Scan(&x.Id, &x.Title, &x.Seller, &x.Price,&x.Stock,&x.SoldQuantity); err != nil {
			return nil, errors.NewInternalServerError("error when tying to gett cashier"+ err.Error())
		}
		results = append(results, x)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}

func (item *Item) List() ([]*Item, *errors.RestErr) {
	stmt, err := datasources.Client.Prepare(queryListItems)
	if err != nil {
		return nil, errors.NewInternalServerError("error when tying to get item" + err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query() //update with limit and skip
	if err != nil {
		return nil, errors.NewInternalServerError("error when tying to get cashier" + err.Error())
	}
	defer rows.Close()

	results := make([]*Item, 0)
	for rows.Next() {
		x := &Item{}
		if err := rows.Scan(&x.Id, &x.Title, &x.Seller, &x.Price,&x.Stock,&x.SoldQuantity); err != nil {
			return nil, errors.NewInternalServerError("error when tying to gett cashier"+ err.Error())
		}
		results = append(results, x)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
	}
	return results, nil	
}


func (item *Item) Delete() *errors.RestErr {
	stmt, err := datasources.Client.Prepare(queryDeleteItem)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get item " + err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(item.Id); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}

func (item *Item) Update() *errors.RestErr {
	stmt, err := datasources.Client.Prepare(queryUpdateItem)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get item " + err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(
		item.Title,
		item.Seller,
		item.Price,
		item.Stock,
		item.SoldQuantity,
		item.Id,
	); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (item *Item) GetById() *errors.RestErr {
	stmt, err := datasources.Client.Prepare(queryById)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get item " + err.Error())
	}
	defer stmt.Close()

	res := stmt.QueryRow(item.Id)

	if err := res.Scan(
		&item.Title,
		&item.Seller,
		&item.Price,
		&item.Stock,
		&item.SoldQuantity,
		&item.CreatedAt,
		); err != nil {
		
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewBadRequestError("no item with this id exists")
		}
		return errors.NewInternalServerError("error reading item into the json" + err.Error())
	}

	return nil
}

func(item *Item) Create() *errors.RestErr {

	stmt, err := datasources.Client.Prepare(queryInsertItem)
	if err != nil {
		return errors.NewInternalServerError("error when trying to get category" + err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(item.Title, item.Seller, item.Price, item.Stock, item.SoldQuantity, item.CreatedAt)
	if saveErr != nil {
		return errors.NewInternalServerError("error when tying to save user"+err.Error())
	}

	itemId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when tying to save user"+err.Error())
	}
	item.Id = itemId

	return nil
}

// func(p *Product) Delete() rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryDeleteProduct)
// 	if err != nil {
// 		logger.Error("error when trying to prepare get cashier statement", err)
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	_, err = stmt.Exec(p.Id)
// 	if err != nil {
// 		logger.Error("error when trying to update user", err)
// 		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
// 	}
// 	return nil
// }

// func(p *Product) Update() rest_errors.RestErr {

// 	temp := &Product{
// 		Id : p.Id,
// 	}
// 	temp.GetById()

// 	if p.CategoryId != 0 {
// 		temp.CategoryId = p.CategoryId
// 	}
// 	if len(p.Name) != 0 {
// 		temp.Name = p.Name
// 	}
// 	if len(p.Image) != 0 {
// 		temp.Image = p.Image
// 	}
// 	if p.Price != 0 {
// 		temp.Price = p.Price
// 	}
// 	if p.Stock != 0 {
// 		temp.Stock = p.Stock
// 	}


// 	stmt, err := cashiers_db.Client.Prepare(queryUpdateProduct)
// 	if err != nil {
// 		logger.Error("error when trying to prepare get cashier statement", err)
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(
// 		temp.CategoryId,
// 		temp.Name,
// 		temp.Image,
// 		temp.Price,
// 		temp.Stock,
// 		temp.Id,
// 	)
// 	if err != nil {
// 		logger.Error("error when trying to update user", err)
// 		return rest_errors.NewInternalServerError("error when tying to update user", errors.New("database error"))
// 	}
// 	return nil
// }

// func(p *Product) GetById() rest_errors.RestErr {
// 	stmt, err := cashiers_db.Client.Prepare(queryById)
// 	if err != nil {
// 		logger.Error("error when trying to prepare get cashier statement", err)
// 		return rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	result := stmt.QueryRow(p.Id)

// 	if err := result.Scan(&p.Id,
// 		&p.CategoryId,
// 		&p.Name,
// 		&p.Image,
// 		&p.Price,
// 		&p.Stock,
// 		&p.UpdatedAt,
// 		&p.CreatedAt,
// 		&p.Discount.Qty, 
// 		&p.Discount.Type,
// 		&p.Discount.Result,
// 		&p.Discount.ExpiredAt,
// 		); err != nil {
// 		logger.Error("error when scan cashier row into cashier struct", err)
// 		return rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
// 	}

// 	return nil
// }

// func(product *Product) Create() rest_errors.RestErr {

// 	product.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z")
// 	product.UpdatedAt = product.CreatedAt

// 	stmt, err := cashiers_db.Client.Prepare(queryInsertProduct)
// 	if err != nil {
// 		logger.Error("error when trying to prepare prepare product create statement", err)
// 		return rest_errors.NewInternalServerError("error when trying to get category", errors.New("database error"))
// 	}
// 	defer stmt.Close()

// 	insertResult, saveErr := stmt.Exec(product.CategoryId, product.Name, product.Image, product.Price, product.Stock, product.Discount.Qty, product.Discount.Type, product.Discount.Result, product.Discount.ExpiredAt, product.UpdatedAt, product.CreatedAt)
// 	if saveErr != nil {
// 		logger.Error("error when trying to save product", saveErr)
// 		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
// 	}

// 	productId, err := insertResult.LastInsertId()
// 	if err != nil {
// 		logger.Error("error when trying to get last insert id after creating a new user", err)
// 		return rest_errors.NewInternalServerError("error when tying to save user", errors.New("database error"))
// 	}
// 	product.Id = productId

// 	return nil
// }

// func (p *Product) List() ([]Product, rest_errors.RestErr) {
// 	stmt, err := cashiers_db.Client.Prepare(queryListProducts)
// 	if err != nil {
// 		logger.Error("error when trying to prepare get cashier statement", err)
// 		return nil, rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query() //update with limit and skip
// 	if err != nil {
// 		logger.Error("error when trying list cahisers", err)
// 		return nil, rest_errors.NewInternalServerError("error when tying to get cashier", errors.New("database error"))
// 	}
// 	defer rows.Close()

// 	results := make([]Product, 0)
// 	for rows.Next() {
// 		var temp Product
// 		if err := rows.Scan(&temp.Id,
// 			&temp.CategoryId,
// 			&temp.Name,
// 			&temp.Image,
// 			&temp.Price,
// 			&temp.Stock,
// 			&temp.UpdatedAt,
// 			&temp.CreatedAt,
// 			&temp.Discount.Qty, 
// 			&temp.Discount.Type,
// 			&temp.Discount.Result,
// 			&temp.Discount.ExpiredAt,
// 			); err != nil {
// 			logger.Error("error when scan cashier row into cashier struct", err)
// 			return nil, rest_errors.NewInternalServerError("error when tying to gett cashier", errors.New("database error"))
// 		}
// 		results = append(results, temp)
// 	}
// 	if len(results) == 0 {
// 		return nil, rest_errors.NewNotFoundError(fmt.Sprintf("no cashiers matching status %s", "ok"))
// 	}
// 	return results, nil	
// }