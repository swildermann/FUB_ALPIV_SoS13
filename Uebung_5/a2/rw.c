#include <pthread.h>
#include <stdio.h>
#include <stdbool.h>

unsigned int rc = 0;
unsigned int rq = 0;
unsigned int wq = 0;
bool busy = false;
pthread_mutex_t mutex = 
	PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t OKtowrite = 
	PTHREAD_COND_INITIALIZER;
pthread_cond_t OKtoread = 
	PTHREAD_COND_INITIALIZER;

void *reader_start() {
	pthread_mutex_lock(&mutex);
	if (busy || wq != 0) {
		pthread_cond_wait(&OKtoread,
		  &mutex);
	}
	rc ++;
	pthread_cond_signal(&OKtoread);
	pthread_mutex_unlock(&mutex);
}
void *reader_end() {
	pthread_mutex_lock(&mutex);
	rc--;
	if (rc ==  0) {
		pthread_cond_signal(&OKtowrite);
	}
	pthread_mutex_unlock(&mutex);
}
void *writer_start() {
	pthread_mutex_lock(&mutex);
	if ( rc != 0 || busy) {
		pthread_cond_wait(&OKtowrite, &mutex);
	}
	busy = true;
	pthread_mutex_unlock(&mutex);
}

void *writer_end() {
	pthread_mutex_lock(&mutex);
	busy = false;
	if ( rq != 0) {
		pthread_cond_signal(&OKtoread);
	} else {
		pthread_cond_signal(&OKtowrite);
	}
	pthread_mutex_unlock(&mutex);
}

main () {
	pthread_t p, q;
	void *f (void *x) {
		reader_start();
		return NULL;
	}
	void *g(void *y)  {
		writer_start();
		return NULL;
	}
	pthread_create(&p, NULL, &f, NULL);
	pthread_create(&q, NULL, &g, NULL);
	(void) pthread_join(p,NULL);
	(void) pthread_join(q,NULL);
}
