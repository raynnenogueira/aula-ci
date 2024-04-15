package main
import "testing"

func TestAdd(t *testing T) { // Função pra testar a função de soma
  result := Add(2, 3)
  expected := 5
  if result != expected {
    t.Errorf("Add(2, 3) = %d", result, expected)
  }
}
