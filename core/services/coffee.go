package services

import (
	"context"
	"log"

	"github.com/FadyGamilM/cfe_api/core/domain"
)

// i created this embedding type around my domain type so i can achieve the architecture separation because in Go i can't create a method on a non-local type which is a type declared in another package such as the domain.Coffee
type Coffee struct {
	domain.Coffee
}

func (c *Coffee) GetAllCoffees() ([]domain.Coffee, error) {
	/*
		1. define a context with time out = 3 seconds to limit my database operations
		2. call cancel on this context to release all used resources for this context
		3. define a query
		4. execute the query using db.QueryWityContext() method and pass the context
		5. handle if there is an error due to the query execution
		6. map the result of the query to my domain core model using the rows.Next and rows.Scan methods
		7. return the result
	*/
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeOut)
	defer cancel()

	query := `SELECT id, name, region, roast_level, price, grind_unit, created_at, updated_at FROM coffees`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Fatalf("cannot execute the GetAllCoffees query, \n \t ERROR : %v \n", err)
	}

	defer rows.Close()

	var coffees []domain.Coffee

	for rows.Next() {
		var c Coffee
		err := rows.Scan(
			&c.Coffee.ID,
			&c.Coffee.Name,
			&c.Coffee.Region,
			&c.Coffee.RoastLevel,
			&c.Coffee.Price,
			&c.Coffee.GrindUnit,
			&c.Coffee.CreatedAt,
			&c.Coffee.UpdatedAt,
		)
		if err != nil {
			log.Fatalf("cannot map the query result into coffee type, \n \t ERROR : %v \n", err)
		} else {
			coffees = append(coffees, c.Coffee)
		}
	}

	return coffees, nil
}
