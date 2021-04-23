pipeline {
    agent {
        node { 
            label "master"
        }   
    } 
    environment {
        GOOS = 'linux'
        GOARCH = "amd64"
        MYPROJECTNAME = 'myddns'
        BUILDBIN = "${MYPROJECTNAME}_${GOARCH}"
        AliyunAccessId = ""
        AliyunAccessKey = ""
        DDNSDomain = ""
    }
    tools {
        go 'go'
    }
    stages {
        stage('Checkout') {
            steps {
                checkout([
                 $class: 'GitSCM',
                    branches: [[name: 'refs/tags/*']],
                    userRemoteConfigs: [[
                        url: 'git@github.com:huyinghuan/ddns.git',
                        credentialsId: '1d6fabe3-4e94-419b-9a48-a90fd3a0f327',
                    ]]
                ])
            }
        }
        stage("编译"){
            environment {
                ProjectVersion = """${sh(
                    returnStdout: true,
                    script: "git describe --tags"
                ).trim()}"""
                BuildTime = """${sh(
                    returnStdout: true,
                    script: "date '+%Y-%m-%d %H:%M:%S'"
                ).trim()}
                """
                argsVersion = "main.Version=${ProjectVersion}"
                argsBuildTime = "main.BuildTime=${BuildTime}"
            }
            steps {
                sh '''
                    go build -o ${BUILDBIN} -mod=vendor -ldflags="-X '${argsVersion}' -X '${argsBuildTime}'"
                '''
                // 存储编译文件
                stash name: "bin_${BUILDBIN}", includes: "${BUILDBIN}"
                sh "rm ${BUILDBIN}"
            }     
        }
        stage("部署"){
            agent { 
                node { 
                    label "LAN-UB"
                    customWorkspace "/home/hyh/service/"
                }
            }
            steps{
                sh "mkdir -p /home/hyh/service"
                // 取出编译二进制文件
                unstash name: "bin_${BUILDBIN}"
                sh """
                    if [ "\$(pm2 id ${MYPROJECTNAME})" != "[]" ]; then 
                        pm2 stop ${MYPROJECTNAME};
                    fi
                    chmod +x ${BUILDBIN}
                    cp ${BUILDBIN} ${MYPROJECTNAME}
                    if [ "\$(pm2 id ${MYPROJECTNAME})" = "[]" ]; then 
                        pm2 start ${MYPROJECTNAME} --name ${MYPROJECTNAME} -l start.${MYPROJECTNAME}.log --interpreter none -- --accessId ${AliyunAccessId}  --accessKey ${AliyunAccessKey} --domain ${DDNSDomain}; 
                    else
                        pm2 restart ${MYPROJECTNAME};
                    fi
                    #等待2秒，判断是否启动成功
                    sleep 2s
                    if [ "\$(pm2 pid ${MYPROJECTNAME})" = "0" ]; then 
                        pm2 delete ${MYPROJECTNAME}
                        echo "启动失败"
                        exit 1; 
                    fi
                    pm2 save
                """
                
            }
        }
    }
}