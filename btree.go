package main

/* 1ページのデータ数上限の半分 1ページのデータ数はMから2M  */
const M = 2
const (
	FALSE = iota
	TRUE
)

type keytype int
type page struct {
	/* データ数 */
	n int
	/* キー  */
	key [M * 2]keytype
	/* 他ページへのポインタ */
	branch [M*2 + 1]*page
}
type pageptr *page

var root page = nil
var key keytype
var done, delted, undersize bool
var newp page
var message string

func newpage() pageptr {
	var p pageptr
	p = new(page)
	return p
}

func search() {
	var p pageptr
	var k int

	p = root
	for p != nil {
		k = 0
		for k < p.n && p.key[k] < key {
			if k < p.n && p.key[k] == key {
				message = "見つかりました"
				return
			}
			k++
		}
		p = p.branch[k]
	}
	message = "見つかりません"
}

func insertitem(p pageptr, k int) {
	var i int

	for i = p.n; i > k; i-- {
		p.key[i] = p.key[i-1]
		p.branch[i+1] = p.branch[i]
	}
	p.key[k] = key
	p.branch[k+1] = newp
	p.n++
}

func split(p pageptr, k int) {
	var j, m int
	var q pageptr
	if k <= M {
		m = M
	} else {
		m = M + 1
	}
	q = newpage()
	for j = m + 1; j <= 2*M; j++ {
		q.key[j-m-1] = p.key[j-1]
		q.branch[j-m] = p.branch[j]
	}
	q.n = 2*M - m
	p.n = m
	if k <= M {
		insertitem(p, k)
	} else {
		insertitem(q, k-m)
	}
	key = p.key[p.n-1]
	q.branch[0] = p.branch[p.n]
	newp = q
}

func insertsub(p pageptr) {
	var k int

	if p == nil {
		done = false
		newp = nil
	}
	k = 0
	for k < p.n && p.key[k] < key {
		k++
	}
	if k < p.n && p.key[k] == key {
		message = "already regist"
		done = true
		return
	}
	insertsub(p.branch[k])
	if done {
		return
	}
	if p.n < 2*M {
		insertitem(p, k)
		done = true
	} else {
		split(p, k)
		done = false
	}
}

func insert() {
	var p pageptr

	message = "registed"
	insertsub(root)
	if done {
		return
	}
	p = newpage()
	p.n = 1
	p.key[0] = key
	p.branch[0] = root
	p.branch[1] = newp
	root = p
}

func printtree(p pageptr) {
	var depth int = 0
	var k int

	if p == nil {
		fmt.Printf(".")
		return
	}
	for k = 0; k < p.n; k++ {
		printtree(p.branch[k])
		fmt.Printf("%d", p.key[k])
	}
	printtree(p.branch[k])
	fmt.Printf(")")
	depth--
}

func main() {
	var s string

	for {
		fmt.Printf("Insert:In, search:Sn (n is num) \n")
		if n, err := fmt.Scanf("%s%d", &s, &key); err == nil {
			fmt.Printf("error: %d\n", n)
			os.Exit(1)
		}
		LBL
		switch key {
		case "I":
			insert()
		case "S":
			printtree()
		default:
			break LBL
		}
	}
}
