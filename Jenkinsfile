node {
    stage('checkout'){
        checkout scm
    }
	stage('Docker build'){
       golang=docker.build('cmpute/testapp')
		}
    stage('Docker push'){
        docker.withRegistry('https://556080742458.dkr.ecr.us-east-1.amazonaws.com/cmpute/testapp', 'ecr:us-east-1:dev-credentials') {
            docker.image('cmpute/testapp').push('latest')
            }
        }
	}

