package models

import "database/sql"

type Problem struct {
	Id           int    `json:"id"`
	Leetcode_num int    `json:"leetcode_num"`
	Problem_name string `json:"problem_name"`
	Description  string `json:"description"`
	Approach_id  int    `json:"approach_id"`
	Pseudocode   string `json:"pseudocode"`
	Diff_id      int    `json:"diff_id"`
}

func GetProblems(db *sql.DB) ([]Problem, error) {
	rows, err := db.Query(
		"SELECT id, leetcode_num, problem_name, description, approach_id, pseudocode, diff_id FROM problems")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	problems := []Problem{}

	for rows.Next() {
		var problem Problem
		if err := rows.Scan(&problem.Id, &problem.Leetcode_num, &problem.Problem_name, &problem.Description,
			&problem.Approach_id, &problem.Pseudocode, &problem.Diff_id); err != nil {
			return nil, err
		}
		problems = append(problems, problem)
	}

	return problems, nil
}
