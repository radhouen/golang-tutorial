package main

import "testing"

func TestArticle_String(t *testing.T) {
	type fields struct {
		Title  string
		Author string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Article{
				Title:  tt.fields.Title,
				Author: tt.fields.Author,
			}
			if got := a.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Circle{
				radius: tt.fields.radius,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Circle{
				radius: tt.fields.radius,
			}
			if got := r.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySqlConnection_Log(t *testing.T) {
	type fields struct {
		dataBaseName     string
		dataBaseUrl      string
		dataBaseUserName string
		dataBasePassword string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Connection := MySqlConnection{
				dataBaseName:     tt.fields.dataBaseName,
				dataBaseUrl:      tt.fields.dataBaseUrl,
				dataBaseUserName: tt.fields.dataBaseUserName,
				dataBasePassword: tt.fields.dataBasePassword,
			}
			if got := Connection.Log(); got != tt.want {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Perimeter(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSShConnection_Log(t *testing.T) {
	type fields struct {
		userName string
		url      string
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Connection := SShConnection{
				userName: tt.fields.userName,
				url:      tt.fields.url,
				password: tt.fields.password,
			}
			if got := Connection.Log(); got != tt.want {
				t.Errorf("Log() = %v, want %v", got, tt.want)
			}
		})
	}
}
