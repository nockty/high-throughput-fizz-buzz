package main

import (
	"io"
	"os"
	"runtime"
)

const (
	BUFFER_SIZE  = 2_000_000
	BUFFER_COUNT = 10

	// maximum number of bytes filled during a 15-lines step
	MAX_FILL_SIZE    = 6*5 + 1*9 + 8*20
	STEPS_PER_BUFFER = BUFFER_SIZE / MAX_FILL_SIZE
)

func main() {
	fizzBuzz(os.Stdout, 8_000_000_000)
}

func fizzBuzz(w io.Writer, n int) {
	workerCount := runtime.NumCPU()
	if workerCount > 1 {
		// let another CPU run the write goroutine
		workerCount -= 1
	}
	// workers consume from tasks
	tasks := make(chan *fillTask, 100)
	for i := 0; i < workerCount; i++ {
		go fillWorker(tasks)
	}

	// write goroutine reads from ordered queue and sends buffers back
	queue := make(chan *fillTask, 100)
	buffers := make(chan []byte, BUFFER_COUNT)
	for i := 0; i < BUFFER_COUNT; i++ {
		buffers <- make([]byte, BUFFER_SIZE)
	}
	done := make(chan struct{})
	go write(w, queue, buffers, done)

	i := 1
	for i+STEPS_PER_BUFFER*15+15 <= n {
		buf := <-buffers

		task := &fillTask{
			buf:   buf,
			start: i,
			done:  make(chan struct{}, 1),
		}
		// the worker will fill numbers until that one
		i += STEPS_PER_BUFFER * 15

		// send task to parallel workers
		tasks <- task
		// send task to ordered write queue
		queue <- task
	}
	close(tasks)

	// Last iteration
	buf := <-buffers
	offset := 0
	// Intermediate buffer for writing integers. The max int64 is 19 digits in base 10.
	var a [19]byte
	for i <= n {
		if i%15 == 0 {
			fillFizzBuzz(buf, offset)
			offset += 9
		} else if i%3 == 0 {
			fillFizz(buf, offset)
			offset += 5
		} else if i%5 == 0 {
			fillBuzz(buf, offset)
			offset += 5
		} else {
			n := fillInt(i, buf, offset, &a)
			offset += n
		}
		i++
	}
	finalTask := &fillTask{
		buf:  buf[:offset],
		done: make(chan struct{}, 1),
	}
	finalTask.done <- struct{}{}
	queue <- finalTask

	close(queue)

	<-done
}

func write(w io.Writer, queue <-chan *fillTask, buffers chan<- []byte, done chan struct{}) {
	for task := range queue {
		<-task.done
		w.Write(task.buf)
		task.buf = task.buf[:BUFFER_SIZE]
		buffers <- task.buf
	}
	done <- struct{}{}
}

type fillTask struct {
	buf   []byte
	start int
	done  chan struct{}
}

func fillWorker(tasks <-chan *fillTask) {
	// Intermediate buffer for writing integers. The max int64 is 19 digits in base 10.
	var a [19]byte

	for task := range tasks {
		offset := 0
		i := task.start
		for k := 0; k < STEPS_PER_BUFFER; k++ {
			i, offset = fillStep(i, task.buf, offset, &a)
		}
		task.buf = task.buf[:offset]
		task.done <- struct{}{}
	}
}

func fillStep(i int, buf []byte, offset int, a *[19]byte) (int, int) {
	n := fillInt(i, buf, offset, a)
	offset += n
	i++
	n = fillInt(i, buf, offset, a)
	offset += n
	fillFizz(buf, offset)
	offset += 5
	i += 2
	n = fillInt(i, buf, offset, a)
	offset += n
	fillBuzz(buf, offset)
	offset += 5
	fillFizz(buf, offset)
	offset += 5
	i += 3
	n = fillInt(i, buf, offset, a)
	offset += n
	i++
	n = fillInt(i, buf, offset, a)
	offset += n
	fillFizz(buf, offset)
	offset += 5
	fillBuzz(buf, offset)
	offset += 5
	i += 3
	n = fillInt(i, buf, offset, a)
	offset += n
	fillFizz(buf, offset)
	offset += 5
	i += 2
	n = fillInt(i, buf, offset, a)
	offset += n
	i++
	n = fillInt(i, buf, offset, a)
	offset += n
	fillFizzBuzz(buf, offset)
	offset += 9
	i += 2
	return i, offset
}

func fillFizz(buf []byte, offset int) {
	buf[offset] = 'F'
	buf[offset+1] = 'i'
	buf[offset+2] = 'z'
	buf[offset+3] = 'z'
	buf[offset+4] = '\n'
}

func fillBuzz(buf []byte, offset int) {
	buf[offset] = 'B'
	buf[offset+1] = 'u'
	buf[offset+2] = 'z'
	buf[offset+3] = 'z'
	buf[offset+4] = '\n'
}

func fillFizzBuzz(buf []byte, offset int) {
	buf[offset] = 'F'
	buf[offset+1] = 'i'
	buf[offset+2] = 'z'
	buf[offset+3] = 'z'
	buf[offset+4] = 'B'
	buf[offset+5] = 'u'
	buf[offset+6] = 'z'
	buf[offset+7] = 'z'
	buf[offset+8] = '\n'
}

// fillInt writes u in base 10 in buf, using a as an intermediate buffer
func fillInt(u int, buf []byte, offset int, a *[19]byte) int {
	i := 19

	for u >= 1000 {
		is := u % 1000 * 3
		u /= 1000
		i -= 3
		a[i+2] = smallsString[is+2]
		a[i+1] = smallsString[is+1]
		a[i+0] = smallsString[is+0]
	}

	// us < 1000
	is := u * 3
	i--
	a[i] = smallsString[is+2]
	if u >= 10 {
		i--
		a[i] = smallsString[is+1]
	}
	if u >= 100 {
		i--
		a[i] = smallsString[is]
	}

	n := 0
	for j := i; j < 19; j++ {
		buf[offset+n] = a[j]
		n++
	}
	buf[offset+n] = '\n'
	return n + 1
}

const smallsString = "000001002003004005006007008009010011012013014015016017018019020021022023024025026027028029030031032033034035036037038039040041042043044045046047048049050051052053054055056057058059060061062063064065066067068069070071072073074075076077078079080081082083084085086087088089090091092093094095096097098099" +
	"100101102103104105106107108109110111112113114115116117118119120121122123124125126127128129130131132133134135136137138139140141142143144145146147148149150151152153154155156157158159160161162163164165166167168169170171172173174175176177178179180181182183184185186187188189190191192193194195196197198199" +
	"200201202203204205206207208209210211212213214215216217218219220221222223224225226227228229230231232233234235236237238239240241242243244245246247248249250251252253254255256257258259260261262263264265266267268269270271272273274275276277278279280281282283284285286287288289290291292293294295296297298299" +
	"300301302303304305306307308309310311312313314315316317318319320321322323324325326327328329330331332333334335336337338339340341342343344345346347348349350351352353354355356357358359360361362363364365366367368369370371372373374375376377378379380381382383384385386387388389390391392393394395396397398399" +
	"400401402403404405406407408409410411412413414415416417418419420421422423424425426427428429430431432433434435436437438439440441442443444445446447448449450451452453454455456457458459460461462463464465466467468469470471472473474475476477478479480481482483484485486487488489490491492493494495496497498499" +
	"500501502503504505506507508509510511512513514515516517518519520521522523524525526527528529530531532533534535536537538539540541542543544545546547548549550551552553554555556557558559560561562563564565566567568569570571572573574575576577578579580581582583584585586587588589590591592593594595596597598599" +
	"600601602603604605606607608609610611612613614615616617618619620621622623624625626627628629630631632633634635636637638639640641642643644645646647648649650651652653654655656657658659660661662663664665666667668669670671672673674675676677678679680681682683684685686687688689690691692693694695696697698699" +
	"700701702703704705706707708709710711712713714715716717718719720721722723724725726727728729730731732733734735736737738739740741742743744745746747748749750751752753754755756757758759760761762763764765766767768769770771772773774775776777778779780781782783784785786787788789790791792793794795796797798799" +
	"800801802803804805806807808809810811812813814815816817818819820821822823824825826827828829830831832833834835836837838839840841842843844845846847848849850851852853854855856857858859860861862863864865866867868869870871872873874875876877878879880881882883884885886887888889890891892893894895896897898899" +
	"900901902903904905906907908909910911912913914915916917918919920921922923924925926927928929930931932933934935936937938939940941942943944945946947948949950951952953954955956957958959960961962963964965966967968969970971972973974975976977978979980981982983984985986987988989990991992993994995996997998999"
