FROM splunk/splunk:9.0.3

USER root
WORKDIR /opt/splunk

RUN touch etc/.ui_login

# ENV SPLUNK_APPS_URL=/tmp/apps/encode-decode-data-for-splunk_105.tgz,/tmp/apps/splunk-add-on-for-unix-and-linux_850.tgz,/tmp/apps/graphviz_12.tgz,/tmp/apps/splunk-app-for-lookup-file-editing_360.tgz,/tmp/apps/splunk-dashboard-examples_825.tgz,/tmp/apps/splunk-common-information-model-cim_502.tgz,/tmp/apps/splunk-timeline-custom-visualization_161.tgz

# COPY --chown=splunk:splunk \
#     third_party_apps/ \
#     /tmp/apps

COPY --chown=splunk:splunk \
    assets/application.css \
    etc/apps/search/appserver/static/application.css

COPY --chown=splunk:splunk \
    assets/forensibus.jpg \
    etc/apps/search/appserver/static/forensibus.jpg

COPY --chown=splunk:splunk \
    config/default.yml \
    /tmp/defaults/default.yml

COPY --chown=splunk:splunk \
    config/local/ \
    etc/system/local

COPY --chown=splunk:splunk \
    config/local/user-prefs.conf \
    etc/users/forensibus/user-prefs/local/user-prefs.conf

COPY --chown=splunk:splunk \
    app/ \
    etc/apps/forensibus/

USER ansible