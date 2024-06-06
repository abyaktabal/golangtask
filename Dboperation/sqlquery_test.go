package dboperation

import (
	"database/sql"
	"reflect"
	"testing"

	structure "server.go/Structure"
)

func TestInsertemployee(t *testing.T) {
	type args struct {
		data structure.Employee
		conn *sql.DB
	}
	tests := []struct {
		name    string
        args    args
        data    structure.Employee
        want    structure.Respose
        wantErr bool
	}{
		// TODO: Add test cases.
		{
            name: "valid case",
            args: args{
                data: structure.Employee{Id: 1, Name: "maya", Position: "ing", Salary: 89}},
            want:    structure.Respose{StatusCode: 0, StatusDesc: "Emoplyee_Id :1 is successfully inserted ."},
            wantErr: false,
        },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Insertemployee(tt.args.data, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("Insertemployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Insertemployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteEmployeeDetais(t *testing.T) {
	type args struct {
		data structure.Respose
		conn *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    structure.Respose
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteEmployeeDetais(tt.args.data, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteEmployeeDetais() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeleteEmployeeDetais() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateEmployeeDetais(t *testing.T) {
	type args struct {
		data structure.EmployeeUpdatedetails
		conn *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    structure.Respose
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateEmployeeDetais(tt.args.data, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateEmployeeDetais() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateEmployeeDetais() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchEmployee(t *testing.T) {
	type args struct {
		data structure.EmployeeSelect
		conn *sql.DB
	}
	tests := []struct {
		name    string
		args    args
		want    structure.Respose
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchEmployee(tt.args.data, tt.args.conn)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}
