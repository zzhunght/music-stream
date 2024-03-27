package utils

import "github.com/jackc/pgx/v5/pgtype"

func ConvertStringToText(input string) pgtype.Text {

	return pgtype.Text{
		String: input,
		Valid:  true,
	}
}
