pipeline {
    agent any
    environment {
        APP_NAME     = "test-app"
        TARGET_NODES = "adminlintas@10.192.15.4 adminlintas@10.192.15.5"
        
        // Memanggil semua secret dari Jenkins Store
        DB_H = credentials('DB_HOST_SECRET')
        DB_U = credentials('DB_USER_SECRET')
        DB_N = credentials('DB_NAME_SECRET')
        DB_P = credentials('DB_PASS_SECRET')
    }
    stages {
        stage('Build & Export') {
            steps {
                sh "docker build -t ${APP_NAME}:local ."
                sh "docker save ${APP_NAME}:local -o /tmp/${APP_NAME}.tar"
            }
        }
        stage('Secure Remote Deploy') {
            steps {
                script {
                    def nodes = TARGET_NODES.split(' ')
                    for (node in nodes) {
                        sh "scp /tmp/${APP_NAME}.tar docker-compose.yml ${node}:~/test-app/"
                        
                        // Menyuntikkan semua variable ke dalam perintah docker-compose up
                        sh """
                        ssh -o StrictHostKeyChecking=no ${node} "
                            cd ~/test-app && \
                            docker load -i ${APP_NAME}.tar && \
                            DB_HOST='${DB_H}' \
                            DB_USER='${DB_U}' \
                            DB_NAME='${DB_N}' \
                            DB_PASSWORD='${DB_P}' \
                            docker-compose up -d && \
                            rm ${APP_NAME}.tar
                        "
                        """
                    }
                }
            }
        }
    }
}
