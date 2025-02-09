package main
import ("testing"
	"time"
	"fmt"
	"github.com/dm1254/pokedexcli/internal/pokecache"
)
func TestCleanInput(t *testing.T){
	cases := []struct{
		input string
		expected []string
	}{
{	
			input : " hello world ",
			expected: []string{"hello","world"},
		},
		{
			input : "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}
	for _, c := range cases{
		actual := cleanInput(c.input)
		actual_len := len(actual)
		expected_len := len(c.expected)
		if actual_len != expected_len{
			t.Errorf(`----------------
			lengths do not match 
			Actual: %d
			Expected: %d`, actual_len, expected_len)
			return 
		}
		for i := range actual{
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf(`-------------
				words do not match
				Actual: %s
				Expected: %s`, word, expectedWord)
				return
			}
		}
	}

}

func TestNewCache(t *testing.T){
	interval := time.Second * 5
	cache := pokecache.NewCache(interval)
	
	if cache.Cache == nil{
		t.Errorf("cache was not properly initialized")
		return 
	}
	if cache.Interval != interval{
		t.Errorf("time Interval was not properly set, got %v, want %v", cache.Interval, interval)
		return 
	}


}

func TestAddGet(t *testing.T){
	const interval = time.Second * 5
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://Example.com",
			val: []byte("test data"),
		},
		{
			key: "http://newexample.com",
			val: []byte("moretestdata"),
		},

	}
	for i, c := range cases{
		t.Run(fmt.Sprintf("Test case %v",i), func(t *testing.T){
			cache := pokecache.NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key); 
			if !ok{
				t.Errorf("Expected key not found")
				return
			}
			if string(val) != string(c.val){
				t.Errorf("Expected value not found")
				return 
			}

		})
	}	

}

func TestReapLoop(t *testing.T){
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := pokecache.NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))
	_, ok := cache.Get("https://example.com"); 
	if !ok{
		t.Errorf("expected to find key")
		return 	
	}
	
	time.Sleep(waitTime)
	
	_,ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return 	
	}

}
