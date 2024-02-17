package utils

import "github.com/jackc/pgx/v5/pgtype"

func ConvertStringToText(input string) pgtype.Text {
	var text pgtype.Text
	text.String = input
	return text
}
