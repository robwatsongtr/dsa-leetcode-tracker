package models

import (
	"database/sql"
)

type Problem struct {
	Id            int    `json:"id"`
	Leetcode_num  int    `json:"leetcode_num"`
	Problem_name  string `json:"problem_name"`
	Description   string `json:"description"`
	Approach_id   int    `json:"approach_id"`
	Approach_name string `json:"approach_name"`
	Pseudocode    string `json:"pseudocode"`
	Diff_id       int    `json:"diff_id"`
	Diff_level    string `json:"diff_level"`
	Category_id   string `json:"category_id"`
	Category_name string `json:"category_name"`
}

// type Category struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

func GetProblems(db *sql.DB) ([]Problem, error) {
	rows, err := db.Query(
		`SELECT 
			problems.id, 
			problems.leetcode_num, 
			problems.problem_name, 
			problems.description, 
			problems.approach_id, 
			approach.approach_name,
			problems.pseudocode, 
			problems.diff_id, 
			difficulty.diff_level,
			COALESCE(categories.id, 0) AS category_id,
			COALESCE(categories.category_name, '') AS category_name
		FROM problems
		INNER JOIN approach ON problems.approach_id = approach.id
		INNER JOIN difficulty ON problems.diff_id = difficulty.id 
		LEFT JOIN problem_categories ON problems.id = problem_categories.problem_id
		LEFT JOIN categories ON problem_categories.category_id = categories.id 
		ORDER BY problems.id `,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	problems := []Problem{}

	for rows.Next() {
		var problem Problem
		if err := rows.Scan(
			&problem.Id,
			&problem.Leetcode_num,
			&problem.Problem_name,
			&problem.Description,
			&problem.Approach_id,
			&problem.Approach_name,
			&problem.Pseudocode,
			&problem.Diff_id,
			&problem.Diff_level,
			&problem.Category_id,
			&problem.Category_name,
		); err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}

	return problems, nil
}
