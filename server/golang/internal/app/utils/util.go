package utils

import "github.com/jackc/pgx/v5/pgtype"

func ConvertStringToText(input string) pgtype.Text {

	return pgtype.Text{
		String: input,
		Valid:  true,
	}
}

func IntToPGType(num int32) pgtype.Int4 {
	return pgtype.Int4{
		Int32: num,
		Valid: true,
	}
}
